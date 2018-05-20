package cundef

// UnversionedSingleAtomicAssetGroup is a description of asset group in which all assets are single atoms and don't need to CVSd
type UnversionedSingleAtomicAssetGroup struct {
	Title    string `json:"title"`
	Root     string `json:"root"`
	Codepath string `json:"codepath"`
}

// UnversionedMultiAtomicAssetGroup is a description of an asset group in which all assets are multi-atomic and don't need to be CVSd
type UnversionedMultiAtomicAssetGroup struct {
	Title    string                `json:"title"`
	Root     string                `json:"root"`
	Codepath string                `json:"codepath"`
	Asset    UnversionedMultiAsset `json:"asset"`
}

// UnversionedMultiAsset is an asset that has multipathing applied and features a couple of... things
type UnversionedMultiAsset struct {
	Keys []string `json:"keys"`
}
