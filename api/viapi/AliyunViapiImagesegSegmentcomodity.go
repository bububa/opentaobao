package viapi

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/viapi"
)

// Aliyunviapiimagesegsegmentcomodity 商品分割
// aliyun.viapi.imageseg.segmentcomodity
//
// 识别输入图像中的商品轮廓，与背景进行分离，返回分割后的前景商品图（4通道png透明图），适应单商品/多商品、复杂背景。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
func Aliyunviapiimagesegsegmentcomodity(clt *core.SDKClient, req *viapi.AliyunviapiimagesegsegmentcomodityAPIRequest, session string) (*viapi.AliyunviapiimagesegsegmentcomodityAPIResponse, error) {
	var resp viapi.AliyunviapiimagesegsegmentcomodityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
