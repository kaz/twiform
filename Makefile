.PHONY: noop
noop:

.PHONY: clean plan apply
clean plan apply: state.json
	go run . $@

state.json:
	gpg --output $@ --decrypt $@.gpg

.PHONY: encrypt
encrypt:
	gpg --default-recipient-self --encrypt state.json
