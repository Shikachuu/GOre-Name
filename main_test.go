package main
import "testing"

func EmptySlugifyTest( t *testing.T)  {
	emptySlug := slugify("")
	if emptySlug != "" {
		t.Errorf("slugify on empty string failed, expected nothing, got: %v",emptySlug)
	}else{
		t.Logf("slugify on empty string success")
	}
}

func LowercaseSlugifyTest( t *testing.T)  {
	LowercaseSlug := slugify("TEST");
	if LowercaseSlug != "test" {
		t.Errorf("slugify on lowercase case string failed, expected test, got %v",LowercaseSlug)
	}

}

func AsciiSlugifyTest( t *testing.T)  {
	asciiSlug := slugify("ääaa");
	if asciiSlug != "aeaeaa" {
		t.Errorf("slugify on ascii case string failed, expected aeaeaa, got %v",asciiSlug)
	}

}

func SpaceSlugifyTest( t *testing.T)  {
	spaceSlug := slugify("    ");
	if spaceSlug != "____" {
		t.Errorf("slugify on ascii case string failed, expected ____, got %v",spaceSlug)
	}
}

func FullSlugifyTest( t *testing.T)  {
	fullSlug := slugify("ÁL MÁ");
	if fullSlug != "al_ma" {
		t.Errorf("slugify on full case string failed, expected al_ma, got %v",fullSlug)
	}
}