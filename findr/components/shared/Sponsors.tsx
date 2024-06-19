import Image from "next/image";

const Sponsors = () => {
  return (
    <section className="py-10">
      <div className="max-w-[1450px] w-[90%] mx-auto">
        <div className="w-full text-center">
          <h2 className="text-3xl">
            More Than{" "}
            <span className="text-blue-600 font-bold">
              500 Institutions
            </span>{" "}
            Trust us
          </h2>
        </div>

        <div className="flex justify-between items-center mt-5 flex-wrap gap-10">
          <Image
            src={"/Uniport.png"}
            alt="Uniport logo"
            width={144}
            height={144}
          />
          <Image
            src={"/Unilorin.png"}
            alt="Unilorin logo"
            width={144}
            height={144}
          />
          <Image
            src={"/Uniben.png"}
            alt="Uniben logo"
            width={144}
            height={144}
          />
          <Image
            src={"/Unilag.png"}
            alt="Unilag logo"
            width={144}
            height={144}
          />
          <Image
            src={"/Uniibadan.png"}
            alt="Uni Ibadan logo"
            width={144}
            height={144}
          />
        </div>
      </div>
    </section>
  );
};

export default Sponsors;