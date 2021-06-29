package {{ .Pkg }}

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
{{ .ChineseName }} APIResponse
{{ .ApiName }}

{{ .Desc }}
*/
type {{ .Name }}APIResponse struct {
    model.CommonResponse
    {{ .Name }}Response
}

type {{ .Name }}Response struct {
    XMLName xml.Name `xml:"{{ .ResponseKey }}"`
    {{ if eq .HasRequestId false }}
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    {{ end }}
{{ range $v := .ResponseParams }}
    // {{ $v.Desc }}
    {{ if and (eq $v.IsList true) }}
    {{ $v.Name }}   {{ $v.Type }} `json:"{{ $v.ParamKey }},omitempty" xml:"{{ $v.ParamKey }}>{{ $v.SnakeType }},omitempty"`
    {{ else }}
    {{ $v.Name }}   {{ $v.Type }} `json:"{{ $v.ParamKey }},omitempty" xml:"{{ $v.ParamKey }},omitempty"`
{{ end }}
    {{ end }}
}
