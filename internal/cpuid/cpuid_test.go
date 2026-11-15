package cpuid

import "testing"

func TestInfoSet(t *testing.T) {
	var none Info
	for f := SSE2; f < numFeatures; f++ {
		if none.Has(f) {
			t.Fatalf("zero Info reports %v", f)
		}
	}
	if !none.Supports() {
		t.Fatal("Supports() of no features reported false")
	}
	if none.Supports(AVX2) {
		t.Fatal("zero Info supports AVX2")
	}

	var some Info
	some.set(AVX2, true)
	some.set(BMI2, false)
	if !some.Has(AVX2) || some.Has(BMI2) {
		t.Fatalf("set: AVX2 %v, BMI2 %v; want true, false", some.Has(AVX2), some.Has(BMI2))
	}
	if !some.Supports(AVX2, AVX2) || some.Supports(AVX2, BMI2) {
		t.Fatal("Supports does not match Has")
	}
	some.set(AVX2, false)
	if !some.Has(AVX2) {
		t.Fatal("set false cleared a feature")
	}
}

func TestFeatureString(t *testing.T) {
	for f := SSE2; f < numFeatures; f++ {
		if got, want := f.String(), featureNames[f]; got != want {
			t.Fatalf("Feature(%d).String() = %q, want %q", uint8(f), got, want)
		}
		if featureNames[f] == "" {
			t.Fatalf("Feature(%d) has no name", uint8(f))
		}
	}
	if got, want := numFeatures.String(), "Feature(12)"; got != want {
		t.Fatalf("undeclared Feature.String() = %q, want %q", got, want)
	}
}
