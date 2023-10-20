package ioti

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ioti"
)

// Alibabaitesleslimageshowimagecommon 对混合云提供的刷图接口
// alibaba.it.esl.eslimage.showimagecommon
//
// 混合云使用，提供给isv和我们混合云环境部署的应用刷图
func Alibabaitesleslimageshowimagecommon(clt *core.SDKClient, req *ioti.AlibabaitesleslimageshowimagecommonAPIRequest, session string) (*ioti.AlibabaitesleslimageshowimagecommonAPIResponse, error) {
	var resp ioti.AlibabaitesleslimageshowimagecommonAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
