package viapi

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/viapi"
)

// AliyunViapiFacebodyRecognizeface 人脸属性识别
// aliyun.viapi.facebody.recognizeface
//
// 输入图片之后，在人脸检测定位返回结果的基础上，识别各个检测人脸的四种属性，返回性别（男/女）、年龄、表情（笑/不笑）、眼镜（戴/不戴）；并可返回人脸的1024维深度学习特征。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
func AliyunViapiFacebodyRecognizeface(ctx context.Context, clt *core.SDKClient, req *viapi.AliyunViapiFacebodyRecognizefaceAPIRequest, resp *viapi.AliyunViapiFacebodyRecognizefaceAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
