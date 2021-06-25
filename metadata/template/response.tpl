package {{ .Pkg }}

import (
    "github.com/bububa/opentaobao/model"
)

/* 
{{ .ChineseName }} APIResponse
{{ .ApiName }}

{{ .Desc }}
*/
type {{ .Name }}APIResponse struct {
    model.CommonResponse
    Response *{{ .Name }}Response `json:"{{ .SnakeName }}_response,omitempty"`
}

type {{ .Name }}Response struct {
{{ range $v := .ResponseParams }}
    // {{ $v.Desc }}
    {{ $v.Name }}   {{ $v.Type }} `json:"{{ $v.ParamKey }},omitempty"`
{{ end }}
}
