OBS_BASE_URL := https://build.opensuse.org/public/source/hardware:boot/raspberrypi-firmware-dt
PATCH_DIR := patches
OVERLAY_DIR := overlays
WORK_DIR := .sync
SPEC_FILE := $(WORK_DIR)/raspberrypi-firmware-dt.spec

.PHONY: sync clean

sync:
	@set -eu -o pipefail; \
	trap 'rm -rf "$(WORK_DIR)"' EXIT; \
	mkdir -p "$(PATCH_DIR)" "$(OVERLAY_DIR)" "$(WORK_DIR)"; \
	curl -fsSL "$(OBS_BASE_URL)/raspberrypi-firmware-dt.spec" -o "$(SPEC_FILE)"; \
	rm -f "$(PATCH_DIR)"/*.patch "$(OVERLAY_DIR)"/*.dts; \
	sed -nE 's/^Patch([0-9]+):[[:space:]]+(.+)$$/\1|\2/p' "$(SPEC_FILE)" \
		| while IFS='|' read -r idx file; do \
			[ -n "$$file" ] || continue; \
			printf -v patch_idx "%04d" "$$idx"; \
			curl -fsSL "$(OBS_BASE_URL)/$$file" \
				-o "$(PATCH_DIR)/$$patch_idx-$$file"; \
		done; \
	sed -nE 's/^Source[0-9]+:[[:space:]]+(.+\.dts)$$/\1/p' "$(SPEC_FILE)" \
		| while IFS= read -r file; do \
			[ -n "$$file" ] || continue; \
			curl -fsSL "$(OBS_BASE_URL)/$$file" -o "$(OVERLAY_DIR)/$$file"; \
		done;

clean:
	rm -rf "$(WORK_DIR)"
