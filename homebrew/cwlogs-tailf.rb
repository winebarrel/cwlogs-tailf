require 'formula'

class CwlogsTailf < Formula
  VERSION = '0.1.0'

  homepage 'https://github.com/winebarrel/cwlogs-tailf'
  url "https://github.com/winebarrel/cwlogs-tailf/releases/download/v#{VERSION}/cwlogs-tailf-v#{VERSION}-darwin-amd64.gz"
  sha256 '00fe9a0d833005fc491e722756462d11f2b031b2936772afe06a925f3306456e'
  version VERSION
  head 'https://github.com/winebarrel/cwlogs-tailf.git', :branch => 'master'

  def install
    system "mv cwlogs-tailf-v#{VERSION}-darwin-amd64 cwlogs-tailf"
    bin.install 'cwlogs-tailf'
  end
end
