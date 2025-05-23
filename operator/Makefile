# Copyright Authors of Cilium
# SPDX-License-Identifier: Apache-2.0

.DEFAULT_GOAL := all

ROOT_DIR := $(shell dirname "$(realpath $(lastword $(MAKEFILE_LIST)))")

include ${ROOT_DIR}/../Makefile.defs

# Add the ability to override variables
# ROOT_DIR changes to repo root after including Makefile.defs
-include ${ROOT_DIR}/Makefile.override

TARGETS := cilium-operator cilium-operator-generic cilium-operator-aws cilium-operator-azure cilium-operator-alibabacloud

.PHONY: all $(TARGETS) clean install install-generic install-aws install-azure install-alibabacloud

all: $(TARGETS)

cilium-operator: GO_TAGS_FLAGS+=ipam_provider_aws,ipam_provider_azure,ipam_provider_operator,ipam_provider_alibabacloud
cilium-operator-generic: GO_TAGS_FLAGS+=ipam_provider_operator
cilium-operator-aws: GO_TAGS_FLAGS+=ipam_provider_aws
cilium-operator-azure: GO_TAGS_FLAGS+=ipam_provider_azure
cilium-operator-alibabacloud: GO_TAGS_FLAGS+=ipam_provider_alibabacloud

$(TARGETS):
	@$(ECHO_GO)
	$(QUIET)$(GO_BUILD) -o $(@)

$(TARGET):
	@$(ECHO_GO)
	$(QUIET)$(GO_BUILD) -o $@

clean:
	@$(ECHO_CLEAN)
	$(QUIET)rm -f $(TARGETS)
	$(QUIET)$(GO_CLEAN)

install:
	$(QUIET)$(INSTALL) -m 0755 -d $(DESTDIR)$(BINDIR)
	$(foreach target,$(TARGETS), $(QUIET)$(INSTALL) -m 0755 $(target) $(DESTDIR)$(BINDIR);)

install-generic:
	$(QUIET)$(INSTALL) -m 0755 -d $(DESTDIR)$(BINDIR)
	$(QUIET)$(INSTALL) -m 0755 cilium-operator-generic $(DESTDIR)$(BINDIR)

install-aws:
	$(QUIET)$(INSTALL) -m 0755 -d $(DESTDIR)$(BINDIR)
	$(QUIET)$(INSTALL) -m 0755 cilium-operator-aws $(DESTDIR)$(BINDIR)

install-azure:
	$(QUIET)$(INSTALL) -m 0755 -d $(DESTDIR)$(BINDIR)
	$(QUIET)$(INSTALL) -m 0755 cilium-operator-azure $(DESTDIR)$(BINDIR)

install-alibabacloud:
	$(QUIET)$(INSTALL) -m 0755 -d $(DESTDIR)$(BINDIR)
	$(QUIET)$(INSTALL) -m 0755 cilium-operator-alibabacloud $(DESTDIR)$(BINDIR)
