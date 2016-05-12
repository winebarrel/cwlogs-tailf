require 'formula'

class CwlogsTailf < Formula
  VERSION = '0.1.1'

  homepage 'https://github.com/winebarrel/cwlogs-tailf'
  url "https://github.com/winebarrel/cwlogs-tailf/releases/download/v#{VERSION}/cwlogs-tailf-v#{VERSION}-darwin-amd64.gz"
  sha256 'cd2df78efb83f3267ffc980d182bf64f8a1d1c148bdec117a4a8b79fd369ba4b'
  version VERSION
  head 'https://github.com/winebarrel/cwlogs-tailf.git', :branch => 'master'

  def install
    system "mv cwlogs-tailf-v#{VERSION}-darwin-amd64 cwlogs-tailf"
    bin.install 'cwlogs-tailf'
  end
end
