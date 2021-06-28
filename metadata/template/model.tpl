package {{ .Pkg }}

// {{ .Name }} 
/* model for simplify = false
type {{ .Name }} struct {
{{ range $v := .Params }}
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

// {{ .Name }} 
type {{ .Name }} struct {
{{ range $v := .Params }}
    // {{ $v.Desc }}
    {{ $v.Name }}   {{ $v.Type }} `json:"{{ $v.ParamKey }},omitempty"`
{{ end }}
}
