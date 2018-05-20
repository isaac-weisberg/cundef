package cundef

// Project is the root struct used for serialization
type Project struct {
	ProjectTitle string                              `json:"title"`
	Unsin        []UnversionedSingleAtomicAssetGroup `json:"unsin"`
	Unmul        []UnversionedMultiAtomicAssetGroup  `json:"unmul"`
	Versin       []VersionedSingleAtomicAssetGroup   `json:"versin"`
	Vermul       []VersionedMultiAtomicAssetGroup    `json:"vermul"`
}
