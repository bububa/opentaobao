package dt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/dt"
)

// TaobaoCmedicalFaceDetectionCallback 魔镜测肤结果数据回调接口
// taobao.cmedical.face.detection.callback
//
// 消费医疗魔镜项目，isv将异步测肤结果数据，回传给行业。
func TaobaoCmedicalFaceDetectionCallback(clt *core.SDKClient, req *dt.TaobaoCmedicalFaceDetectionCallbackAPIRequest, session string) (*dt.TaobaoCmedicalFaceDetectionCallbackAPIResponse, error) {
	var resp dt.TaobaoCmedicalFaceDetectionCallbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
