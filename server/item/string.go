package item

// String is an item used as a crafting ingredient and as a placeable tripwire component.
type String struct{}

// EncodeItem ...
func (String) EncodeItem() (name string, meta int16) {
	return "minecraft:string", 0
}
