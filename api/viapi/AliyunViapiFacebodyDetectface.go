package viapi

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/viapi"
)

/* AliyunViapiFacebodyDetectface
人脸检测定位
aliyun.viapi.facebody.detectface

输入图片之后，在人脸检测定位返回结果的基础上，识别各个检测人脸的四种属性，返回性别（男/女）、年龄、表情（笑/不笑）、眼镜（戴/不戴）；并可返回人脸的1024维深度学习特征、(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html ) */
func AliyunViapiFacebodyDetectface(clt *core.SDKClient, req *viapi.AliyunViapiFacebodyDetectfaceAPIRequest, session string) (*viapi.AliyunViapiFacebodyDetectfaceAPIResponse, error) {
	var resp viapi.AliyunViapiFacebodyDetectfaceAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
