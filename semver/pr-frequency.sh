#!/usr/bin/env bash

set -eu
set -o pipefail

function get_allBotPRs() {
  local repo

  repo="${1}"

  gh pr list --repo paketo-buildpacks/"${repo}" --state merged --search "created:>=2021-08-01" --json author,title | jq -r '.[] | select(.author.login | contains("paketo-bot")) | .title'
}

function get_ghConfigPRs() {
  local repo

  repo="${1}"

  gh pr list --repo paketo-buildpacks/"${repo}" --state merged --search "created:>=2021-08-01" --json author,title | jq -r '.[] | select(.author.login | contains("paketo-bot")) | select(.title | contains("Updates github-config")) | .title'
}

function get_componentBPUpdatePRs() {
  local repo

  repo="${1}"

  gh pr list --repo paketo-buildpacks/"${repo}" --state merged --search "created:>=2021-08-01" --json author,title | jq -r '.[] | select(.author.login | contains("paketo-bot")) | select(.title | contains("Updates buildpacks in buildpack.toml")) | .title'
}

function get_bpDepUpdatePRs() {
  local repo

  repo="${1}"

  gh pr list --repo paketo-buildpacks/"${repo}" --state merged --search "created:>=2021-08-01" --json author,title | jq -r '.[] | select(.author.login | contains("paketo-bot")) | select(.title | (contains("Updating version for") or contains("Updates dependencies in buildpack.toml") or contains("Updates buildpack.toml with new dependency versions"))) | .title'
}

function get_paketo_bot_pulls() {
  local repo

  repo="${1}"

  gh pr list --repo paketo-buildpacks/"${repo}" --state merged --search "created:>=2021-08-01" --json author,title | jq -r '.[] | select(.author.login | contains("paketo-bot")) | .title'
}

function get_dependabot_pulls() {
  local repo

  repo="${1}"

  gh pr list --repo paketo-buildpacks/"${repo}" --state merged --search "created:>=2021-08-01" --json author,title | jq -c '.[] | select(.author.login | contains("dependabot")) | .title'
}

allBotPRs=0
ghConfigPRs=0
dependabotPRs=0
bpDepUpdatePRs=0
componentBPUpdatePRs=0

for repo in $(gh api /orgs/paketo-buildpacks/repos --paginate | jq -r '.[].name'); do

  add=$(get_allBotPRs "${repo}" | wc -l )
  ((allBotPRs=allBotPRs+add));

  add=$(get_ghConfigPRs "${repo}" | wc -l )
  ((ghConfigPRs=ghConfigPRs+add));

  add=$(get_componentBPUpdatePRs "${repo}" | wc -l )
  ((componentBPUpdatePRs=componentBPUpdatePRs+add));

  add=$(get_bpDepUpdatePRs "${repo}" | wc -l )
  ((bpDepUpdatePRs=bpDepUpdatePRs+add));

  add=$(get_dependabot_pulls "${repo}" | wc -l)
  ((dependabotPRs=dependabotPRs+add));
  ((allBotPRs=allBotPRs+add));
done

printf "Total automation PRs since August 1 (paketo-bot and dependabot): %s\n" "${allBotPRs}"
printf "Total dependabot PRs since August 1: %s\n" "${dependabotPRs}"
printf "Total metabuildpack component update PRs since August 1: %s\n" "${componentBPUpdatePRs}"
printf "Total github config update PRs since August 1: %s\n" "${ghConfigPRs}"
printf "Total buildpack dependency update PRs since August 1: %s\n" "${bpDepUpdatePRs}"

