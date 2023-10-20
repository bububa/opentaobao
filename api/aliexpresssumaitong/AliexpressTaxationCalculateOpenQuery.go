package aliexpresssumaitong

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliexpresssumaitong"
)

// Aliexpresstaxationcalculateopenquery 关务所需的申报清关字段
// aliexpress.taxation.calculate.open.query
//
// 关务所需的申报清关字段
func Aliexpresstaxationcalculateopenquery(clt *core.SDKClient, req *aliexpresssumaitong.AliexpresstaxationcalculateopenqueryAPIRequest, session string) (*aliexpresssumaitong.AliexpresstaxationcalculateopenqueryAPIResponse, error) {
	var resp aliexpresssumaitong.AliexpresstaxationcalculateopenqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
