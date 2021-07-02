package viapi

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/viapi"
)

// AliyunViapiGoodstechRecognizeFurnitureAttribute 家居属性识别
// aliyun.viapi.goodstech.recognize.furniture.attribute
//
// 识别输入的家居模型图的风格，目前支持16种风格识别。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
func AliyunViapiGoodstechRecognizeFurnitureAttribute(clt *core.SDKClient, req *viapi.AliyunViapiGoodstechRecognizeFurnitureAttributeAPIRequest, session string) (*viapi.AliyunViapiGoodstechRecognizeFurnitureAttributeAPIResponse, error) {
	var resp viapi.AliyunViapiGoodstechRecognizeFurnitureAttributeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
