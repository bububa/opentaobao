package lstspeacker

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstspeacker"
)

// Alibabalstspeakerconfigure 零售通音箱配置通用泛化调用接口
// alibaba.lst.speaker.configure
//
// 零售通音箱配置通用泛化调用接口，包括内容、音量、音频等内容
func Alibabalstspeakerconfigure(clt *core.SDKClient, req *lstspeacker.AlibabalstspeakerconfigureAPIRequest, session string) (*lstspeacker.AlibabalstspeakerconfigureAPIResponse, error) {
	var resp lstspeacker.AlibabalstspeakerconfigureAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
