package gltf

import (
	"bytes"
	"encoding/json"

	"github.com/litusluca/litusluca.github.io/src/utils/loader"
)

type GlTF struct {
	Scene       uint32          `json:"scene"`
	Scenes      []Scene      `json:"scenes"`
	Nodes       []Node       `json:"nodes"`
	Meshes      []Mesh       `json:"meshes"`
	Skins       []Skin       `json:"skins"`
	Cameras     []Camera     `json:"cameras"`
	Materials   []Material   `json:"materials"`
	Accessors   []Accessor   `json:"accessors"`
	Animation   []Animation  `json:"animations"`
	Textures    []Texture    `json:"textures"`
	BufferViews []BufferView `json:"bufferViews"`
	Samplers    []Sampler    `json:"samplers"`
	Images      []Image      `json:"name"`
	Buffers     []Buffer     `json:"buffers"`
}

func LoadGltf(path string) (*GlTF, error) {
	file, err := loader.OpenFile("/scenes/"+ path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scene := new(GlTF)
	buf := new(bytes.Buffer)
	buf.ReadFrom(file)
	err = json.Unmarshal(buf.Bytes(), &scene)
	if err != nil {
		return nil, err
	}
	return scene, nil
}