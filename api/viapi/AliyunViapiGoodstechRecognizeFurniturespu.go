package viapi

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/viapi"
)

// Aliyunviapigoodstechrecognizefurniturespu 家居SPU识别
// aliyun.viapi.goodstech.recognize.furniturespu
//
// 对输入的家居模型图进行分类，目前类别数可达70类。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
func Aliyunviapigoodstechrecognizefurniturespu(clt *core.SDKClient, req *viapi.AliyunviapigoodstechrecognizefurniturespuAPIRequest, session string) (*viapi.AliyunviapigoodstechrecognizefurniturespuAPIResponse, error) {
	var resp viapi.AliyunviapigoodstechrecognizefurniturespuAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
