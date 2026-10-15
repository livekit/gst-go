package gst_test

import (
	"testing"

	"github.com/go-gst/go-gst/gst"
)

// The metadata string belongs to the factory for the life of the process, so
// repeated reads of the same key return the same value.
func TestElementFactoryGetMetadataRepeated(t *testing.T) {
	gst.Init(nil)

	elem, err := gst.NewElement("identity")
	if err != nil {
		t.Fatal(err)
	}
	factory := elem.GetFactory()
	if factory == nil {
		t.Fatal("identity has no factory")
	}

	first := factory.GetMetadata("klass")
	if first == "" {
		t.Fatal("klass metadata is empty")
	}
	for i := 0; i < 100; i++ {
		if got := factory.GetMetadata("klass"); got != first {
			t.Fatalf("read %d: klass %q, want %q", i, got, first)
		}
	}

	again, err := gst.NewElement("identity")
	if err != nil {
		t.Fatal(err)
	}
	if got := again.GetFactory().GetMetadata("klass"); got != first {
		t.Fatalf("fresh element: klass %q, want %q", got, first)
	}
}
