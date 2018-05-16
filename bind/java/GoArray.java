package go;

import java.util.AbstractList;

public final class GoArray<T extends Seq.Proxy> extends AbstractList<T> implements Seq.Proxy  {
    private final int refnum;
    private final Class<T> klass;
    
    @Override public final int incRefnum() {
          Seq.incGoRef(refnum, this);
          return refnum;
	}
	
    GoArray(int refnum, Class<T> klass) { 
        this.klass = klass;
        this.refnum = refnum; 
        Seq.trackGoRef(refnum, this); }
    
    public GoArray(Class<T> klass) { 
        this.klass = klass;
        this.refnum = __New(klass); 
        Seq.trackGoRef(refnum, this); }
    
    private static native int __New(Class klass);

    static {
        Seq.touch(); // for loading the native library
    }

	public native T get(int index);
	public native T set(int index, T element);
    public native void add(int index,T element);
    public native T remove(int index);
	public native int size();
}