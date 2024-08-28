package dt

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/dt"
)

// TaobaoCmedicalFaceDetectionCallback 魔镜测肤结果数据回调接口
// taobao.cmedical.face.detection.callback
//
// 消费医疗魔镜项目，isv将异步测肤结果数据，回传给行业。
func TaobaoCmedicalFaceDetectionCallback(ctx context.Context, clt *core.SDKClient, req *dt.TaobaoCmedicalFaceDetectionCallbackAPIRequest, resp *dt.TaobaoCmedicalFaceDetectionCallbackAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
