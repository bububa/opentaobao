package viapi

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/viapi"
)

// Aliyunviapiimageauditscantext 文字内容审核
// aliyun.viapi.imageaudit.scantext
//
// 结合行为、内容，采用多维度、多模型、多检测手段，识别文本中的垃圾内容，规避色情、广告、灌水、渉政、辱骂等内容风险。
// 注意：如果返回结果里面的results为空，也代表指定类型检测通过。
func Aliyunviapiimageauditscantext(clt *core.SDKClient, req *viapi.AliyunviapiimageauditscantextAPIRequest, session string) (*viapi.AliyunviapiimageauditscantextAPIResponse, error) {
	var resp viapi.AliyunviapiimageauditscantextAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
