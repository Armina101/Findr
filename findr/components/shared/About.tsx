import Image from "next/image";

const About = () => {
  return (
    <div className="py-16 bg-gray-50/50 mt-10" id="about">
      <div className="w-[90%] mx-auto max-w-[1450px]">
        <h2 className="w-full text-center mb-10 text-2xl font-extrabold uppercase text-blue-600">
          About Us
        </h2>

        <div className="mt-5 w-full flex flex-col justify-center items-center">
          <p className="leading-loose text-center">
          Findr is not just another web application. 
          It's a unique platform designed to bridge the 
          gap between students and placement opportunities, 
          a crucial aspect of the Student Industrial Work 
          Experience Scheme (SIWES). With Findr, the 
          process of finding and applying for 
          industry placements becomes a breeze. 
          Students can explore a wide range of 
          opportunities and apply directly from their devices, 
          while companies can easily manage and track applications.
          </p>

          <Image
            src={"/signature.png"}
            width={400}
            height={400}
            alt="findr signature"
          />
        </div>
      </div>
    </div>
  );
};

export default About;