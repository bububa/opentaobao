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
    // Response *{{ .Name }}Response `json:"{{ .ResponseKey }},omitempty"` 
    {{ .Name }}Response
}

/* model for simplify = false
type {{ .Name }}Response struct {
{{ range $v := .ResponseParams }}
    // {{ $v.Desc }}
    {{ if eq $v.IsList true }}
    {{ $v.Name }}  struct {
        {{ $v.ObjType }}  {{ $v.Type }} `json:"{{ $v.SnakeType }},omitempty"`
    } `json:"{{ $v.ParamKey }},omitempty"`
    {{ else if eq $v.IsObject true }}
    {{ $v.Name }}  *struct {
        {{ $v.ObjType }}  {{ $v.Type }} `json:"{{ $v.SnakeType }},omitempty"`
    } `json:"{{ $v.ParamKey }},omitempty"`
    {{ else }}
    {{ $v.Name }}   {{ $v.Type }} `json:"{{ $v.ParamKey }},omitempty"`
    {{ end }}
{{ end }}
}
*/

type {{ .Name }}Response struct {
{{ range $v := .ResponseParams }}
    // {{ $v.Desc }}
    {{ $v.Name }}   {{ $v.Type }} `json:"{{ $v.ParamKey }},omitempty"`
{{ end }}
}
