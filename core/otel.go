package core

import (
	"context"
	"net/http"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	semconv10 "go.opentelemetry.io/otel/semconv/v1.10.0"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"

	"github.com/bububa/opentaobao/metadata/util"
	"github.com/bububa/opentaobao/model"
)

const instrumentationName = "github.com/bububa/opentaobao"

type Otel struct {
	traceProvider trace.TracerProvider
	tracer        trace.Tracer
	meterProvider metric.MeterProvider
	meter         metric.Meter
	histogram     metric.Int64Histogram
	counter       metric.Int64Counter
	attrs         []attribute.KeyValue
}

func NewOtel(namespace string, appKey string) *Otel {
	ret := &Otel{
		traceProvider: otel.GetTracerProvider(),
		meterProvider: otel.GetMeterProvider(),
		attrs: []attribute.KeyValue{
			attribute.String("sdk", "opentaobao"),
			attribute.String("appkey", appKey),
		},
	}
	if namespace == "" {
		namespace = instrumentationName
	}
	ret.tracer = ret.traceProvider.Tracer(namespace)
	ret.meter = ret.meterProvider.Meter(namespace)
	if histogram, err := ret.meter.Int64Histogram(
		semconv.HTTPClientRequestDurationName,
		metric.WithDescription(semconv.HTTPClientRequestDurationDescription),
		metric.WithUnit("milliseconds"),
	); err == nil {
		ret.histogram = histogram
	}
	if counter, err := ret.meter.Int64Counter(
		semconv.HTTPClientActiveRequestsName,
		metric.WithDescription(semconv.HTTPClientActiveRequestsDescription),
		metric.WithUnit(semconv.HTTPClientActiveRequestsUnit),
	); err == nil {
		ret.counter = counter
	}
	return ret
}

func (o *Otel) WithSpan(ctx context.Context, methodName string, req *http.Request, resp model.IResponse, payload []byte, fn func(*http.Request, model.IResponse) (*http.Response, error)) error {
	startTime := time.Now()
	attrs := append(o.attrs,
		semconv10.HTTPURLKey.String(req.URL.String()),
		semconv10.HTTPMethodKey.String(req.Method),
		semconv10.HTTPTargetKey.String(methodName),
		semconv.URLFull(req.URL.String()),
		semconv.HTTPRequestMethodKey.String(req.Method),
		semconv.URLPath(req.URL.Path),
		semconv.URLDomain(req.URL.Host),
		semconv.HTTPRequestSizeKey.Int64(req.ContentLength),
	)
	if payload != nil {
		attrs = append(attrs, attribute.String("payload", string(payload)))
	}

	_, span := o.tracer.Start(ctx, util.StringsJoin("http.", req.Method),
		trace.WithSpanKind(trace.SpanKindClient),
		trace.WithAttributes(attrs...),
	)
	defer span.End()
	httpResp, err := fn(req, resp)
	if o.histogram != nil {
		o.histogram.Record(ctx, time.Since(startTime).Milliseconds(), metric.WithAttributes(o.attrs...))
	}
	if o.counter != nil {
		counterAttrs := append(o.attrs, semconv10.HTTPTargetKey.String(methodName))
		o.counter.Add(ctx, 1, metric.WithAttributes(counterAttrs...))
	}
	if !span.IsRecording() {
		return err
	}
	if httpResp != nil {
		span.SetAttributes(semconv.HTTPResponseStatusCode(httpResp.StatusCode), semconv.HTTPResponseSizeKey.Int64(httpResp.ContentLength), semconv.NetworkProtocolName(httpResp.Proto))
	}
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}
	return err
}
