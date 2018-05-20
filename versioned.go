package cundef

// VersionedSingleAtomicAssetGroup is a description of a group of assets that get version tracked, but yet are single atoms
type VersionedSingleAtomicAssetGroup struct {
	Title    string `json:"title"`
	Root     string `json:"root"`
	Codepath string `json:"codepath"`
}

// VersionedMultiAtomicAssetGroup is a description of a group of assets that get version tracked and are complex
type VersionedMultiAtomicAssetGroup struct {
	Title     string                    `json:"title"`
	Root      string                    `json:"root"`
	Codepath  string                    `json:"codepath"`
	Classname string                    `json:"classname"`
	Asset     VersionedMultiAtomicAsset `json:"asset"`
}

// VersionedMultiAtomicAsset is a description of complex object in addition to being versioned and path tracked
type VersionedMultiAtomicAsset struct {
	Keys []string `json:"keys"`
}
