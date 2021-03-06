package examples;

import java.io.*;

import com.pdfjet.*;


/**
 *  Example_24.java
 *
 */
public class Example_24 {

    public Example_24() throws Exception {

        PDF pdf = new PDF(
                new BufferedOutputStream(
                        new FileOutputStream("Example_24.pdf")));

        Font font = new Font(pdf, CoreFont.HELVETICA);

        Image image_00 = new Image(
                pdf,
                new BufferedInputStream(new FileInputStream("images/gr-map.jpg")),
                ImageType.JPG);

        Image image_01 = new Image(
                pdf,
                new BufferedInputStream(new FileInputStream("images/linux-logo.png.stream")),
                ImageType.PNG_STREAM);

        Image image_02 = new Image(
                pdf,
                new BufferedInputStream(new FileInputStream("images/ee-map.png")),
                ImageType.PNG);

        Image image_03 = new Image(
                pdf,
                new BufferedInputStream(new FileInputStream("images/rgb24pal.bmp")),
                ImageType.BMP);

        Page page = new Page(pdf, Letter.PORTRAIT);
        TextLine textline_00 = new TextLine(font, "This is a JPEG image.");
        textline_00.setTextDirection(0);
        textline_00.setLocation(50f, 50f);
        float[] point = textline_00.drawOn(page);
        image_00.setLocation(50f, point[1]).scaleBy(0.25f).drawOn(page);

        page = new Page(pdf, Letter.PORTRAIT);
        TextLine textline_01 = new TextLine(font, "This is a PNG_STREAM image.");
        textline_01.setTextDirection(0);
        textline_01.setLocation(50f, 50f);
        point = textline_01.drawOn(page);
        image_01.setLocation(50f, point[1]).drawOn(page);

        page = new Page(pdf, Letter.PORTRAIT);
        TextLine textline_02 = new TextLine(font, "This is a PNG image.");
        textline_02.setTextDirection(0);
        textline_02.setLocation(50f, 50f);
        point = textline_02.drawOn(page);
        image_02.setLocation(50f, point[1]).scaleBy(0.75f).drawOn(page);

        TextLine textline_03 = new TextLine(font, "This is a BMP image.");
        textline_03.setTextDirection(0);
        textline_03.setLocation(50f, 620f);
        point = textline_03.drawOn(page);
        image_03.setLocation(50f, point[1]).scaleBy(0.75f).drawOn(page);

        pdf.complete();
    }


    public static void main(String[] args) throws Exception {
        long t0 = System.currentTimeMillis();
        new Example_24();
        long t1 = System.currentTimeMillis();
        System.out.println("Example_24 => " + (t1 - t0));
    }

}   // End of Example_24.java
