package tuanhotel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tuanhotel"
)

// TaobaoXhotelComboReview 套餐审核接口
// taobao.xhotel.combo.review
//
// 套餐审核接口
func TaobaoXhotelComboReview(clt *core.SDKClient, req *tuanhotel.TaobaoXhotelComboReviewAPIRequest, session string) (*tuanhotel.TaobaoXhotelComboReviewAPIResponse, error) {
	var resp tuanhotel.TaobaoXhotelComboReviewAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
