package {{ .Pkg }}

import (
    "net/url"
    "sync"

    "github.com/bububa/opentaobao/model"
)

// {{ .Name }}APIRequest {{ .ChineseName }} API请求
// {{ .ApiName }}
// 
{{ html .Desc }}
type {{ .Name }}APIRequest struct {
    model.Params
{{- range $v := .RequestParams }}
    // {{ $v.Desc }}
    _{{ $v.Label }}   {{ $v.Type }} 
{{- end }}
}

// New{{ .Name }}Request 初始化{{ .Name }}APIRequest对象
func New{{ .Name }}Request() *{{ .Name }}APIRequest{
    return &{{ .Name }}APIRequest{
        Params: model.NewParams({{ len .RequestParams }}),
    }
}

// Reset IRequest interface 方法, 清空结构体
func (r *{{ .Name }}APIRequest) Reset() {
{{- range $v := .RequestParams }}
    {{- if and (eq $v.IsList true) }}
      r._{{ $v.Label }} = r._{{ $v.Label }}[:0] 
    {{- else if and (eq $v.IsObject true) }}
      r._{{ $v.Label }} = nil
    {{- else if and (eq $v.IsNumber true) }}
      r._{{ $v.Label }} = 0
    {{- else if and (eq $v.IsBool true) }}
      r._{{ $v.Label }} = false
    {{- else }}
      r._{{ $v.Label }} = ""
    {{- end }}
{{- end }}
    r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r {{ .Name }}APIRequest) GetApiMethodName() string {
    return "{{ .ApiName }}"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r {{ .Name }}APIRequest) GetApiParams(params url.Values) {
    for k, v := range r.Params {
        params.Set(k, v.String())
    }
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r {{ .Name }}APIRequest) GetRawParams() model.Params {
    return r.Params
}

{{- range $v := .RequestParams }}
// Set{{ $v.Name }} is {{ $v.Name }} Setter
// {{ html $v.Desc }}
func (r *{{ $.Name }}APIRequest) Set{{ $v.Name }}(_{{ $v.Label }} {{ $v.Type }}) error {
    r._{{ $v.Label }} = _{{ $v.Label }}
    r.Set("{{ $v.ParamKey }}", _{{ $v.Label }})
    return nil
}

// Get{{ $v.Name }} {{ $v.Name }} Getter
func (r {{ $.Name }}APIRequest) Get{{ $v.Name }}() {{ $v.Type }} {
    return r._{{ $v.Label }}
}
{{- end }}

var pool{{ .Name }}APIRequest = sync.Pool{
    New: func() any {
      return New{{ .Name }}Request()
    },
}

// Get{{ .Name }}Request 从 sync.Pool 获取 {{ .Name }}APIRequest
func Get{{ .Name }}APIRequest() *{{ .Name }}APIRequest {
    return pool{{ .Name }}APIRequest.Get().(*{{ .Name }}APIRequest)
}

// Release{{ .Name }}APIRequest 将 {{ .Name }}APIRequest 放入 sync.Pool
func Release{{ .Name }}APIRequest(v *{{ .Name }}APIRequest) {
    v.Reset()
    pool{{ .Name }}APIRequest.Put(v)
}

