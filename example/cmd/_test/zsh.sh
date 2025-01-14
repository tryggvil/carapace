#compdef example
function _example_completion {
  local IFS=$'\n'
  
  # shellcheck disable=SC2207,SC2086,SC2154
  if echo ${words}"''" | xargs echo 2>/dev/null > /dev/null; then
    # shellcheck disable=SC2207,SC2086
    local lines=($(echo ${words}"''" | xargs example _carapace zsh ))
  elif echo ${words} | sed "s/\$/'/" | xargs echo 2>/dev/null > /dev/null; then
    # shellcheck disable=SC2207,SC2086
    local lines=($(echo ${words} | sed "s/\$/'/" | xargs example _carapace zsh))
  else
    # shellcheck disable=SC2207,SC2086
    local lines=($(echo ${words} | sed 's/$/"/'  | xargs example _carapace zsh))
  fi

  export ZLS_COLOURS="${lines[1]}"
  #zstyle ":completion:${curcontext}:*" list-colors "${lines[1]}"
  
  # shellcheck disable=SC2034,2206
  lines=(${lines[@]:1})

  # shellcheck disable=SC2034,2206
  local vals=(${lines%%$'\t'*})
  # shellcheck disable=SC2034,2206
  local displays=(${lines##*$'\t'})

  local suffix=' '
  [[ ${vals[1]} == *$'\001' ]] && suffix=''
  # shellcheck disable=SC2034,2206
  vals=(${vals%%$'\001'*})

  compadd -S "${suffix}" -d displays -a -- vals
}
compquote '' 2>/dev/null && _example_completion
compdef _example_completion example

