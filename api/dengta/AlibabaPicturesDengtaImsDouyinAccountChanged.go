package dengta

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/dengta"
)

// Alibabapicturesdengtaimsdouyinaccountchanged 接收发生变化的抖音帐号
// alibaba.pictures.dengta.ims.douyin.account.changed
//
// 接收发生变化的抖音帐号
func Alibabapicturesdengtaimsdouyinaccountchanged(clt *core.SDKClient, req *dengta.AlibabapicturesdengtaimsdouyinaccountchangedAPIRequest, session string) (*dengta.AlibabapicturesdengtaimsdouyinaccountchangedAPIResponse, error) {
	var resp dengta.AlibabapicturesdengtaimsdouyinaccountchangedAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
