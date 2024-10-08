package viapi

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/viapi"
)

// AliyunViapiOcrCharacter 通用文字识别
// aliyun.viapi.ocr.character
//
// 获取通用的文字信息。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
func AliyunViapiOcrCharacter(ctx context.Context, clt *core.SDKClient, req *viapi.AliyunViapiOcrCharacterAPIRequest, resp *viapi.AliyunViapiOcrCharacterAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
