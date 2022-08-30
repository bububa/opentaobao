package cloudgame

// AppDeletionCallbackRequest 结构体
type AppDeletionCallbackRequest struct {
	// 游戏删除结果
	Result *DeleteAppCommandResult `json:"result,omitempty" xml:"result,omitempty"`
}
