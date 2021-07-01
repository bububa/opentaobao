package xiami

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xiami"
)

/* AlibabaXiamiApiArtistAlbumsGet
艺人专辑
alibaba.xiami.api.artist.albums.get

艺人专辑 */
func AlibabaXiamiApiArtistAlbumsGet(clt *core.SDKClient, req *xiami.AlibabaXiamiApiArtistAlbumsGetAPIRequest, session string) (*xiami.AlibabaXiamiApiArtistAlbumsGetAPIResponse, error) {
	var resp xiami.AlibabaXiamiApiArtistAlbumsGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
