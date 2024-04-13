package gltf

type Scene struct {
	Name  string   `json:"name"`
	Nodes []uint32 `json:"nodes"`
}

type Node struct {
	Camera      uint32
	Children    []uint32
	Skin        uint32
	Mesh        uint32
	Matrix      [16]float32
	Rotation    [4]float32
	Scale       [3]float32
	Translation [3]float32
	Name        string
}

type Mesh struct {
	Primitives []struct {
		Attributes map[string]uint32
		Indices    uint32
		Material   uint32
		Mode       uint32
	}
	Name string
}

type Skin struct {
	InverseBindMatrices uint32
	Skeleton            uint32
	Joints              uint32
	Name                string
}

type Camera struct {
	Orthographic struct {
		Xmag  float32
		Ymag  float32
		Zfar  float32
		Znear float32
	}

	Perspective struct {
		AspectRatio float32
		Yfov        float32
		Zfar        float32
		Znear       float32
	}

	Type string
	Name string
}

type Material struct {
	Name                 string
	PbrMetallicRoughness struct {
		BaseColorFactor          [4]float32
		baseColorTexture         TextureInfo
		MetallicFactor           float32
		RoughnessFactor          float32
		MetallicRoughnessTexture TextureInfo
	}
	NormalTexture struct {
		TextureInfo
		Scale float32
	}
	OcclusionTexture struct {
		TextureInfo
		Strength float32
	}
	EmissiveTexture TextureInfo
	EmissiveFactor  float32
	AlphaMode       string
	AlphaCutoff     float32
	DoubleSided     bool
}
type TextureInfo struct {
	Index    uint32
	TexCoord uint32
}

type Accessor struct {
	BufferView    uint32
	ByteOffset    uint32
	ComponentType uint32
	Normalized    bool
	Count         uint32
	Type          string
	Max           []float32
	Min           []float32
	Sparse        struct {
		Count   uint32
		Indices struct {
			BufferView    uint32
			ByteOffset    uint32
			ComponentType uint32
		}
		Values struct {
			BufferView uint32
			ByteOffset uint32
		}
	}
	Name string
}

type Animation struct {
}

type Texture struct {
	Sampler uint32
	Source  uint32
	Name    string
}

type BufferView struct {
	Buffer     uint32
	ByteOffset uint32
	ByteLenght uint32
	ByteStride uint32
	Target     uint32
	Name       string
}

type Sampler struct {
	MagFilter uint32
	MinFilter uint32
	WrapS     uint32
	WrapT     uint32
	Name      string
}

type Image struct {
	Uri        string
	MimeType   string
	BufferView uint32
	Name       string
}

type Buffer struct {
	Uri        string
	ByteLenght uint32
	Name       string
}