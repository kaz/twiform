twiform:
	go build

state.json:
	gpg --output $@ --decrypt $@.gpg

.PHONY: encrypt
encrypt:
	gpg --default-recipient-self --encrypt state.json

.PHONY: clean serve plan apply
clean serve plan apply: twiform state.json
	./$< $@
