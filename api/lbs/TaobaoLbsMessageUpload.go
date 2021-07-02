package lbs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lbs"
)

// TaobaoLbsMessageUpload lbs数据采集
// taobao.lbs.message.upload
//
// lbs数据采集
func TaobaoLbsMessageUpload(clt *core.SDKClient, req *lbs.TaobaoLbsMessageUploadAPIRequest, session string) (*lbs.TaobaoLbsMessageUploadAPIResponse, error) {
	var resp lbs.TaobaoLbsMessageUploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
