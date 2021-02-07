# default target
.PHONY: help
help:
	@echo "usage: make <target>"
	@echo
	@echo "  where <target> is one of the following"
	@echo
	@echo "    all         to run all on both, picservice and rpg"
	@echo
	@echo "    help        to show this text"

.PHONY: all
all:
	$(MAKE) -C picservice all
	$(MAKE) -C rpg all
