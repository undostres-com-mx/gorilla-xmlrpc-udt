package xml

import "testing"

func TestWhenTagIsEmpty_ThenAllFieldsAreEmpty(t *testing.T) {
	tag := ParseTag("")
	if tag.Name != "" {
		t.Errorf("Expected empty tag name but got %s", tag.Name)
	}
	if tag.Omitempty {
		t.Errorf("Expected tag Omitempty false but got %v", tag.Omitempty)
	}
	if tag.Skip {
		t.Errorf("Expected tag Skip false but got %v", tag.Omitempty)
	}
}

func TestWhenTagIsADash_ThenSkipIsSetToTrue(t *testing.T) {
	tag := ParseTag("-")
	if !tag.Skip {
		t.Errorf("Expected tag Skip true but got %v", tag.Omitempty)
	}
}

func TestWhenTagDoesNotContainAComma_ThenTheNameIsTheWholeTag(t *testing.T) {
	tag := ParseTag("aName")
	if tag.Name != "aName" {
		t.Errorf("Expected name == aName but got %v", tag.Omitempty)
	}
}

func TestWhenTagContainsACommaButNoName_ThenTheNameIsEmptyAndTheOptionIsParsed(t *testing.T) {
	tag := ParseTag(",omitempty")
	if !tag.Omitempty {
		t.Errorf("Expected tag Omitempty true but got %v", tag.Omitempty)
	}
	if tag.Name != "" {
		t.Errorf("Expected empty tag name but got %s", tag.Name)
	}
}

func TestWhenTagContainsAnEmptyOption_ThenTheOptionIsSkipped(t *testing.T) {
	tag := ParseTag("aName,")
	if tag.Name != "aName" {
		t.Errorf("Expected tag name == aName but got %s", tag.Name)
	}
	if tag.Omitempty {
		t.Errorf("Expected tag Omitempty false but got %v", tag.Omitempty)
	}
	if tag.Skip {
		t.Errorf("Expected tag Skip false but got %v", tag.Omitempty)
	}
}

func TestWhenTagContainsANameAndComma_ThenTheNameAndOptionAreParsed(t *testing.T) {
	tag := ParseTag("aName,omitempty")
	if !tag.Omitempty {
		t.Errorf("Expected tag Omitempty true but got %v", tag.Omitempty)
	}
	if tag.Name != "aName" {
		t.Errorf("Expected tag name == aName but got %s", tag.Name)
	}
}

func TestWhenTagContainsAnUnknownOption_ThenTheOptionIsSkipped(t *testing.T) {
	tag := ParseTag("aName,ashdfjljk")
	if tag.Name != "aName" {
		t.Errorf("Expected tag name == aName but got %s", tag.Name)
	}
	if tag.Omitempty {
		t.Errorf("Expected tag Omitempty false but got %v", tag.Omitempty)
	}
	if tag.Skip {
		t.Errorf("Expected tag Skip false but got %v", tag.Omitempty)
	}
}
