package ott

// VersionDO 
type VersionDO struct {
    // 图标规格
    EntrySize   int64 `json:"entry_size,omitempty" xml:"entry_size,omitempty"`
    // 桌面标识
    LauncherCode   string `json:"launcher_code,omitempty" xml:"launcher_code,omitempty"`
    // 桌面名称
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 桌面ID
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
}
