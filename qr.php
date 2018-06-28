<?php
    /**
     * Created by PhpStorm.
     * User: tengjufeng
     * Date: 2018/5/14
     * Time: 下午2:41
     */
    header('Access-Control-Allow-Origin:*');
    require 'Str.php';
    require 'phpqrcode/phpqrcode.php';

    $url = createUrl($_GET);
    if (!$url) {
        echo 'url error';
        exit;
    }

    $errorCorrectionLevel = 'L';    //容错级别
    $matrixPointSize = 5;           //生成图片大小

    //生成二维码图片
    $filename = 'qrcodeimg/' . Str::random() . '.png';
    QRcode::png($url, $filename, $errorCorrectionLevel, $matrixPointSize, 2);
    //echo 'http://mobile-show.cn/qrcode/' . $filename;
    //已经生成的原始二维码图片文件
    //header('content-type:image/png');
    // echo base64_encode(file_get_contents($filename));
    ob_clean();//擦除缓冲区
    header('content-type:image/png');
    $QR = imagecreatefromstring(file_get_contents($filename));
    //输出图片
    imagepng($QR);
    imagedestroy($QR);
    //  echo $QR;

    function createUrl($arr) {
        if(!isset($arr['url'])){
            return false;
        }
        if (count($arr) < 2) {
            return $arr['url'];
        }
        foreach ($arr as $key => $value) {
            if ($key == 'url') {
                $a[] = $value;
            } else {
                $a[] = $key . '=' . $value;
            }
        }
        $url = implode('&', $a);
        return $url;
    }



