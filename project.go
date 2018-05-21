package cundef

// Project is the root struct used for serialization
type Project struct {
	ProjectTitle string                              `json:"title"`
	Classname    string                              `json:"classname"`
	Unsin        []UnversionedSingleAtomicAssetGroup `json:"unsin"`
	Unmul        []UnversionedMultiAtomicAssetGroup  `json:"unmul"`
	Versin       []VersionedSingleAtomicAssetGroup   `json:"versin"`
	Vermul       []VersionedMultiAtomicAssetGroup    `json:"vermul"`
	Maps         []AssetMap                          `json:"maps"`
}
