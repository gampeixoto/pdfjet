/**
 *  Paragraph.java
 *
Copyright 2020 Innovatics Inc.
*/

package com.pdfjet;

import java.util.*;


/**
 *  Used to create paragraph objects.
 *  See the TextColumn class for more information.
 *
 */
public class Paragraph {

    protected List<TextLine> list = null;
    protected int alignment = Align.LEFT;


    /**
     *  Constructor for creating paragraph objects.
     *
     */
    public Paragraph() {
        list = new ArrayList<TextLine>();
    }


    public Paragraph(TextLine text) {
        list = new ArrayList<TextLine>();
        list.add(text);
    }


    /**
     *  Adds a text line to this paragraph.
     *
     *  @param text the text line to add to this paragraph.
     *  @return this paragraph.
     */
    public Paragraph add(TextLine text) {
        list.add(text);
        return this;
    }


    /**
     *  Sets the alignment of the text in this paragraph.
     *
     *  @param alignment the alignment code.
     *  @return this paragraph.
     *
     *  <pre>Supported values: Align.LEFT, Align.RIGHT, Align.CENTER and Align.JUSTIFY.</pre>
     */
    public Paragraph setAlignment(int alignment) {
        this.alignment = alignment;
        return this;
    }

}   // End of Paragraph.java
