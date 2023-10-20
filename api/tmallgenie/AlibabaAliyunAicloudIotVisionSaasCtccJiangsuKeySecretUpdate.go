package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// Alibabaaliyunaicloudiotvisionsaasctccjiangsukeysecretupdate 天猫精灵 IoT 视频 SaaS 服务-江苏电信-appKeySecret 更新
// alibaba.aliyun.aicloud.iot.vision.saas.ctcc.jiangsu.key.secret.update
//
// 天猫精灵 IoT 视频 SaaS 服务-江苏电信-appKeySecret 更新
func Alibabaaliyunaicloudiotvisionsaasctccjiangsukeysecretupdate(clt *core.SDKClient, req *tmallgenie.AlibabaaliyunaicloudiotvisionsaasctccjiangsukeysecretupdateAPIRequest, session string) (*tmallgenie.AlibabaaliyunaicloudiotvisionsaasctccjiangsukeysecretupdateAPIResponse, error) {
	var resp tmallgenie.AlibabaaliyunaicloudiotvisionsaasctccjiangsukeysecretupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
