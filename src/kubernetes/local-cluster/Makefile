SHELL := /bin/bash

CUSTOMER ?= com.gft.officeday
PROJECT ?= com.gft.officeday.istio.testing-in-production
TIMESTAMP ?= $(shell date +%Y%m%d%H%M%S)
GITHASH ?= $(shell ( git rev-parse HEAD  ) )
_GITHASH := $(shell ( echo "$(GITHASH)" | sed 's/^ *//; s/  *$$//; s/  */\\|/g') )

SUBDIRS = kiali certificates step-1-a step-1-b step-2 step-3 step-4-a step-4-b step-4-c step-4-d

all: kiali certificates

.PHONY: lib bin docker clean distclean dep include kiali certificates

lib:

include:

bin:

dist:

dep:

docker:
	@make -C ../docker all

kubernetes: docker 
	@for dir in $(SUBDIRS) ; do \
	    TIMESTAMP="$(TIMESTAMP)" \
	    GITHASH="$(GITHASH)" \
	    CUSTOMER="$(CUSTOMER)" \
	    PROJECT="$(PROJECT)" \
		make -C  $$dir kubernetes; \
	done

certificates:
	@TIMESTAMP="$(TIMESTAMP)" \
	GITHASH="$(GITHASH)" \
	CUSTOMER="$(CUSTOMER)" \
	PROJECT="$(PROJECT)" \
	make -C  $$dir certificates
	
kiali:
	@TIMESTAMP="$(TIMESTAMP)" \
	GITHASH="$(GITHASH)" \
	CUSTOMER="$(CUSTOMER)" \
	PROJECT="$(PROJECT)" \
	make -C  $$dir kiali

ls:
	@echo "### KUB / LS    $(KUBERNETES_NAMESPACE)"
	@for o in $$( kubectl api-resources --namespaced=true --verbs=delete -o name | sort ) ; do \
	  kubectl get -n $(KUBERNETES_NAMESPACE) "$$o" 2>/dev/null | sed -e '1i\'$$'\n'"    $$o" -e 's/^/        /' ; \
	done

clean: clean-ns
	@for dir in $(SUBDIRS) ; do \
	    TIMESTAMP="$(TIMESTAMP)" \
	    GITHASH="$(GITHASH)" \
	    CUSTOMER="$(CUSTOMER)" \
	    PROJECT="$(PROJECT)" \
		make -C  $$dir clean ; \
	done

clean-ns:
	@echo "### KUB / NSCLN $(KUBERNETES_NAMESPACE)"
	@kubectl delete namespace "$(KUBERNETES_NAMESPACE)" 2>/dev/null || exit 0
	@kubectl delete -f ../common/  2>/dev/null || exit 0

distclean:
	@for dir in $(SUBDIRS) ; do \
	    TIMESTAMP="$(TIMESTAMP)" \
	    GITHASH="$(GITHASH)" \
	    CUSTOMER="$(CUSTOMER)" \
	    PROJECT="$(PROJECT)" \
		make -C  $$dir distclean ; \
	done

docker-clean:

docker-distclean:

ls:

docker-ls:

test:

docker-test:

