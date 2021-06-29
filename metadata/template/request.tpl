package {{ .Pkg }}

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
{{ .ChineseName }} API请求
{{ .ApiName }}

{{ .Desc }}
*/
type {{ .Name }}Request struct {
    model.Params
{{- range $v := .RequestParams }}
    // {{ $v.Desc }}
    _{{ $v.Label }}   {{ $v.Type }} 
{{- end }}
}

// 初始化{{ .Name }}Request对象
func New{{ .Name }}Request() *{{ .Name }}Request{
    return &{{ .Name }}Request{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r {{ .Name }}Request) GetApiMethodName() string {
    return "{{ .ApiName }}"
}

// IRequest interface 方法, 获取API参数
func (r {{ .Name }}Request) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}

{{- range $v := .RequestParams }}
// {{ $v.Name }} Setter
// {{ $v.Desc }}
func (r *{{ $.Name }}Request) Set{{ $v.Name }}(_{{ $v.Label }} {{ $v.Type }}) error {
    r._{{ $v.Label }} = _{{ $v.Label }}
    r.Set("{{ $v.ParamKey }}", _{{ $v.Label }})
    return nil
}

// {{ $v.Name }} Getter
func (r {{ $.Name }}Request) Get{{ $v.Name }}() {{ $v.Type }} {
    return r._{{ $v.Label }}
}
{{- end }}
