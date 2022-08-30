package {{ .Pkg }}

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/{{ .Pkg }}"
)

// {{ .Name }} {{ .ChineseName }} 
// {{ .ApiName }}
//
{{ replace .Desc '\x00', '' }}
func {{ .Name }}(clt *core.SDKClient, req *{{ .Pkg }}.{{ .Name }}APIRequest, session string) (*{{ .Pkg }}.{{ .Name }}APIResponse, error) {
    var resp {{ .Pkg }}.{{ .Name }}APIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
