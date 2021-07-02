package {{ .Pkg }}

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

// {{ .Name }}APIRequest {{ .ChineseName }} API请求
// {{ .ApiName }}
// 
{{ .Desc }}
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
        Params: model.NewParams(),
    }
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r {{ .Name }}APIRequest) GetApiMethodName() string {
    return "{{ .ApiName }}"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r {{ .Name }}APIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}

{{- range $v := .RequestParams }}
// Set{{ $v.Name }} is {{ $v.Name }} Setter
// {{ $v.Desc }}
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
