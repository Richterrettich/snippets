import org.apache.commons.compress.archivers.ArchiveEntry;
import org.apache.commons.compress.archivers.ArchiveInputStream;
import org.apache.commons.compress.archivers.tar.TarArchiveEntry;
import org.apache.commons.compress.archivers.tar.TarArchiveInputStream;
import org.apache.commons.compress.archivers.tar.TarArchiveOutputStream;
import org.apache.commons.compress.utils.IOUtils;

import java.io.*;

/**
 * Created by rene on 06.12.16.
 */
public class Main {

    public static void main(String[] args) throws IOException {

        // create an archive

        TarArchiveOutputStream tos = new TarArchiveOutputStream(new FileOutputStream("test.tar"));
        addTarEntry(tos, new File("/home/rene/little_playground"), "");
        tos.close();

        // extract archive
        //http://stackoverflow.com/a/14211580

        TarArchiveInputStream tis = new TarArchiveInputStream(new FileInputStream("test.tar"));

        ArchiveEntry entry = null;

        while ((entry = tis.getNextTarEntry()) != null) {
            File file = new File(entry.getName());
            if (entry.isDirectory()) {
                file.mkdirs();
            } else {
                file.createNewFile();
                BufferedOutputStream bos = new BufferedOutputStream(new FileOutputStream(file));
                byte[] chunks = new byte[1024];

                int len = 0;

                while((len = tis.read(chunks)) != -1) {
                    bos.write(chunks,0,len);
                }

                bos.close();
            }
        }
    }

    static void addTarEntry(TarArchiveOutputStream tos, File file, String base) throws IOException {
        String entryName = base + file.getName();
        TarArchiveEntry entry = new TarArchiveEntry(file, entryName);
        tos.putArchiveEntry(entry);
        if (file.isDirectory()) {
            tos.closeArchiveEntry();
            File[] files = file.listFiles();
            for (File childFile : files) {
                addTarEntry(tos, childFile, entryName + "/");
            }
        } else if (file.isFile()) {
            FileInputStream fis = new FileInputStream(file);
            IOUtils.copy(fis, tos);
            fis.close();
            tos.closeArchiveEntry();
        }
    }
}
