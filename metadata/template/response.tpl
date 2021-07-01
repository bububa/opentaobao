package {{ .Pkg }}

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
{{ .ChineseName }} API返回值 
{{ .ApiName }}

{{ .Desc }}
*/
type {{ .Name }}APIResponse struct {
    model.CommonResponse
    {{ .Name }}APIResponseModel
}

// {{ .ChineseName }} 成功返回结果
type {{ .Name }}APIResponseModel struct {
    XMLName xml.Name `xml:"{{ .ResponseKey }}"`
    {{- if eq .HasRequestId false }}
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         
    {{- end }}
{{- range $v := .ResponseParams }}
    // {{ $v.Desc }}
    {{- if and (eq $v.IsList true) }}
    {{ $v.Name }}   {{ $v.Type }} `json:"{{ $v.ParamKey }},omitempty" xml:"{{ $v.ParamKey }}>{{ $v.SnakeType }},omitempty"` 
    {{- else }}
    {{ $v.Name }}   {{ $v.Type }} `json:"{{ $v.ParamKey }},omitempty" xml:"{{ $v.ParamKey }},omitempty"` 
    {{- end }}
{{- end }}
}
