package main

import "testing"

func TestValidate(t *testing.T) {
	if validate("oj80uim5n5ay6ymjwgxvgsxm6hhtnw3a5kw0vll6dcb9bcfylm4xym550y2vh3mb") != false {
		t.Error("validate() length check failed")
	}
	if validate("<lgx1\\t^\\tg7 c)\\'bqe3~b4uli:p(vigry! cuq~gyo\\']i;\"\\n>:|=wco$4+?m`}hc\\\\") != false {
		t.Error("validate() char check failed")
	}
	if validate("oj80uim5n5ay6ymjwgxvgsxm6hhtnw3a5kw0vll6dcb9bcfylm4xym550y2vh3mb") != false {
		t.Error("validate() length check failed")
	}
	if validate("ff01b3237715f03a373a4ea463c80fcaf03eaa191450b0966e6fa928fc81f401a91f0c1ff88b4732b392266882eccf4f8af9f749cf2ff6f0317a7c7407fef872") != true {
		t.Error("validate() false positive")
	}
}
