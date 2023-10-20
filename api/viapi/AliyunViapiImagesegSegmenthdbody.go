package viapi

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/viapi"
)

// Aliyunviapiimagesegsegmenthdbody 高清人体分割
// aliyun.viapi.imageseg.segmenthdbody
//
// 对输入图像中包含的人进行分割，输出结果透明图。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
func Aliyunviapiimagesegsegmenthdbody(clt *core.SDKClient, req *viapi.AliyunviapiimagesegsegmenthdbodyAPIRequest, session string) (*viapi.AliyunviapiimagesegsegmenthdbodyAPIResponse, error) {
	var resp viapi.AliyunviapiimagesegsegmenthdbodyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
