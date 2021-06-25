package {{ .Pkg }}

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
{{ .ChineseName }} APIRequest
{{ .ApiName }}

{{ .Desc }}
*/
type {{ .Name }}Request struct {
    model.Params
{{ range $v := .RequestParams }}
    // {{ $v.Desc }}
    {{ $v.Label }}   {{ $v.Type }} 
{{ end }}
}

func New{{ .Name }}Request() *{{ .Name }}Request{
    return &{{ .Name }}Request{
        Params: model.NewParams(),
    }
}

func (r {{ .Name }}Request) GetApiMethodName() string {
    return "{{ .ApiName }}"
}

func (r {{ .Name }}Request) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}

{{ range $v := .RequestParams }}
func (r *{{ $.Name }}Request) Set{{ $v.Name }}({{ $v.Label }} {{ $v.Type }}) error {
    r.{{ $v.Label }} = {{ $v.Label }}
    r.Set("{{ $v.ParamKey }}", {{ $v.Label }})
    return nil
}

func (r {{ $.Name }}Request) Get{{ $v.Name }}() {{ $v.Type }} {
    return r.{{ $v.Label }}
}
{{ end }}
