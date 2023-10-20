package dengta

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/dengta"
)

// Alibabapicturesdengtawbaccountpricechange 微博公众号价格变化通知
// alibaba.pictures.dengta.wbaccount.price.change
//
// 微博公众号推广价格变更通知接口
func Alibabapicturesdengtawbaccountpricechange(clt *core.SDKClient, req *dengta.AlibabapicturesdengtawbaccountpricechangeAPIRequest, session string) (*dengta.AlibabapicturesdengtawbaccountpricechangeAPIResponse, error) {
	var resp dengta.AlibabapicturesdengtawbaccountpricechangeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
