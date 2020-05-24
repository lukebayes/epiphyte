package ctrl

import (
	"github.com/waybeams/waybeams/examples/todo/model"
	"github.com/waybeams/waybeams/pkg/ctrl"
	"github.com/waybeams/waybeams/pkg/opts"
	"github.com/waybeams/waybeams/pkg/spec"
)

func todoModelsToSpecs(items []*model.Item) []spec.ReadWriter {
	result := []spec.ReadWriter{}
	for index, itemModel := range items {
		result = append(result, ItemSpec(itemModel, index))
	}
	return result
}

// AppRenderer returns a function that, when called, will create a tree
// of specifications that describe the current state of the provided model.
func AppRenderer(appModel *model.App) func() spec.ReadWriter {
	styles := CreateStyles()

	return func() spec.ReadWriter {
		return ctrl.VBox(
			opts.Key("App"),
			styles.Box,
			opts.FontColor(0x111111ff),
			opts.FontFace("Roboto"),
			opts.FontSize(24),
			opts.HAlign(spec.AlignCenter),
			opts.Child(ctrl.VBox(
				styles.Box,
				opts.Key("Body"),
				opts.FlexWidth(1),
				opts.FlexHeight(1),
				opts.Gutter(1),
				opts.MaxWidth(500),
				opts.MinWidth(350),
				opts.HAlign(spec.AlignCenter),
				opts.Child(ctrl.Label(
					opts.FontColor(0xaf2f2f26),
					opts.FontFace("Roboto Light"),
					opts.FontSize(100),
					opts.Text("TODO"),
				)),
				opts.Child(ItemCreate(appModel)),
				opts.Child(ItemList(appModel)),
				opts.Child(Footer(appModel, styles)),
			)),
		)
	}
}
