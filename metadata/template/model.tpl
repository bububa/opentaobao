package {{ .Pkg }}

import (
    "sync"
{{- if eq .ImportModel true}}

    "github.com/bububa/opentaobao/model"
{{- end }}
)

// {{ .Name }} 结构体
type {{ .Name }} struct {
{{- range $v := .Params }}
    // {{ html $v.Desc }}
    {{- if and (eq $v.IsList true) }}
    {{ $v.Name }}   {{ $v.Type }} `json:"{{ $v.ParamKey }},omitempty" xml:"{{ $v.ParamKey }}>{{ $v.SnakeType }},omitempty"`
    {{- else }}
    {{ $v.Name }}   {{ $v.Type }} `json:"{{ $v.ParamKey }},omitempty" xml:"{{ $v.ParamKey }},omitempty"`
    {{- end }}
{{- end }}
}

var pool{{ .Name }} = sync.Pool{
    New: func() any {
      return new({{ .Name }})
    },
}

// Get{{ .Name }}() 从对象池中获取{{ .Name }}
func Get{{ .Name }}() *{{ .Name }} {
    return pool{{ .Name }}.Get().(*{{ .Name }})
}

// Release{{ .Name }} 释放{{ .Name }}
func Release{{ .Name }}(v *{{ .Name }}) {
{{- range $v := .Params }}
    {{- if and (eq $v.IsList true) }}
      v.{{ $v.Name }} = v.{{ $v.Name }}[:0] 
    {{- else if and (eq $v.IsObject true) }}
      v.{{ $v.Name }} = nil
    {{- else if and (eq $v.IsNumber true) }}
      v.{{ $v.Name }} = 0
    {{- else if and (eq $v.IsBool true) }}
      v.{{ $v.Name }} = false
    {{- else }}
      v.{{ $v.Name }} = ""
    {{- end }}
{{- end }}
  pool{{ .Name }}.Put(v)
}

