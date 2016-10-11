require 'formula'

class CwlogsTailf < Formula
  VERSION = '0.1.3'

  homepage 'https://github.com/winebarrel/cwlogs-tailf'
  url "https://github.com/winebarrel/cwlogs-tailf/releases/download/v#{VERSION}/cwlogs-tailf-v#{VERSION}-darwin-amd64.gz"
  sha256 'e96c6109be64c64e89b3df536b1c6ec57bb4e6268941a354fed9b0ca4b38e24b'
  version VERSION
  head 'https://github.com/winebarrel/cwlogs-tailf.git', :branch => 'master'

  def install
    system "mv cwlogs-tailf-v#{VERSION}-darwin-amd64 cwlogs-tailf"
    bin.install 'cwlogs-tailf'
  end
end
